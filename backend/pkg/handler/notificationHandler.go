package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// NotificationHandler handles HTTP requests related to notifications.
type NotificationHandler struct {
	notificationRepo *repository.NotificationRepository
	sessionRepo      *repository.SessionRepository
	groupMemberRepo *repository.GroupMemberRepository
	groupRepo *repository.GroupRepository
	userRepo *repository.UserRepository
	invitationRepo *repository.InvitationRepository
}

// NewNotificationHandler creates a new instance of NotificationHandler.
// It takes a NotificationRepository and a SessionRepository as parameters.
// Returns a pointer to the newly created NotificationHandler.
func NewNotificationHandler(notificationRepo *repository.NotificationRepository, sessionRepo *repository.SessionRepository, groupMemberRepo *repository.GroupMemberRepository, groupRepo *repository.GroupRepository, userRepo *repository.UserRepository, invitationRepo *repository.InvitationRepository) *NotificationHandler {
	return &NotificationHandler{notificationRepo: notificationRepo, sessionRepo: sessionRepo, groupMemberRepo: groupMemberRepo, groupRepo: groupRepo, userRepo: userRepo, invitationRepo: invitationRepo}
}

func (h *NotificationHandler) CreateNotification(userID, senderID int, messageType, message string) error {
    notification := model.Notification{
        UserId:   userID,
        SenderId: senderID,
        Type:     messageType,
        Message:  message,
        IsRead:   false,
    }
    _, err := h.notificationRepo.CreateNotification(notification)
    return err
}

func (h *NotificationHandler) CreateGroupNotification(userID, groupID int, message string) error {
	notification := model.Notification{
		UserId:  userID,
		GroupId: groupID,
		Type:    "group",
		Message: message,
		IsRead:  false,
	}
	_, err := h.notificationRepo.CreateNotification(notification)
	return err
}

// GetAllNotificationsHandler retrieves all notifications and responds
func (h *NotificationHandler) GetAllNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	notifications, err := h.notificationRepo.GetAllNotifications()
	if err != nil {
		http.Error(w, "Failed to get notifications: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

// GetNotificationByIDHandler retrieves a specific notification by its ID and responds with a JSON object.
func (h *NotificationHandler) GetNotificationByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	notification, err := h.notificationRepo.GetNotificationByID(id)
	if err != nil {
		http.Error(w, "Failed to get notification: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notification)
}

// MarkNotificationAsReadHandler marks a notification as read based on its ID.
func (h *NotificationHandler) MarkNotificationAsReadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	// Mark the notification as read
	err = h.notificationRepo.MarkNotificationAsRead(id)
	if err != nil {
		http.Error(w, "Failed to mark notification as read: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message or appropriate response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification marked as read successfully!"))
}

func (h *NotificationHandler) NotifyGroupDeletion(groupID int) error {
	// Get the list of group members.
    members, err := h.groupMemberRepo.GetGroupMembers(groupID); if err != nil {
        return err
    }
	// Get the group title.
	groupTitle, err := h.groupRepo.GetGroupTitleByID(groupID); if err != nil {
		return err
	}
    // Construct a notification message.
    message := fmt.Sprintf("The group '%s' has been deleted.", groupTitle)

    // Create notifications for each group member.
    for _, member := range members {
        // Create a new notification.
		err = h.CreateNotification(member.UserId, 0, "group", message)
        if err != nil {
            return err
        }
    }

    return nil
}

// notifyGroupAdmin notifies the group admin that a user has requested to join the group.
func (h *NotificationHandler) NotifyGroupAdmin(groupID, userID int) error {
	groupTitle, err := h.groupRepo.GetGroupTitleByID(groupID)
	if err != nil {
		return err
	}

	username, err := h.userRepo.GetUsernameByID(userID)
	if err != nil {
		return err
	}
	message := username + "has requested to join your group:" + groupTitle
	// Get the group admin.
	adminID, err := h.groupMemberRepo.GetGroupAdminByID(groupID)
	if err != nil {
		return err
	}
	return h.CreateNotification(adminID, userID, "group", message)
}

// notifyUserRequestApproved notifies the user that their request was approved.
func (h *NotificationHandler) NotifyUserRequestApproved(userID, groupID int) error {
	message := fmt.Sprintf("Your request to join the group %d has been approved.", groupID)
	return h.CreateNotification(userID, 0, "group", message)
}

// Function to notify the user about the declined request.
func (h *NotificationHandler) NotifyUserDecline(userID, groupID int) error {
	// message
	groupTitle, err := h.groupRepo.GetGroupTitleByID(groupID)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("Your request to join the group %s has been declined.", groupTitle)
	return h.CreateNotification(userID, 0, "group", message)
}

// notifyUserInvitation notifies the user about the group invitation.
func (h *NotificationHandler) NotifyUserInvitation(userID, groupID int) error {
	groupTitle, err := h.groupRepo.GetGroupTitleByID(groupID)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("You have been invited to join the group %s.", groupTitle)
	return h.CreateGroupNotification(userID, groupID, message)
}

// notifyGroupOfNewMember notifies the group members about the new member.
func (h *NotificationHandler) NotifyGroupOfNewMember(groupID, joinUserID int) error {
	group, err := h.groupRepo.GetGroupByID(groupID)
	if err != nil {
		return err
	}

	username, err := h.userRepo.GetUsernameByID(joinUserID)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("'%s' has joined the group '%s'.", username, group.Title)

	members, err := h.groupMemberRepo.GetGroupMembers(groupID)
	if err != nil {
		return err
	}

	for _, member := range members {
		if member.UserId != joinUserID {
			err := h.CreateGroupNotification(member.UserId, groupID, message)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NotifyInvitationDecline notifies the group owner about the declined invitation.
func (h *NotificationHandler) NotifyInvitationDecline(id string) error {
	// Get the invitation details from the repository based on the ID.
	invitation, err := h.invitationRepo.GetGroupInvitationByID(id); if err != nil {
		return err
	}
	username, err := h.userRepo.GetUsernameByID(invitation.JoinUserId); if err != nil {
		return err
	}
	groupTitle, err := h.groupRepo.GetGroupTitleByID(invitation.GroupId); if err != nil {
		return err
	}
	// Construct a notification message.
	message := fmt.Sprintf("The user %s has declined your invitation to join the group %s.", username, groupTitle)

	return h.CreateGroupNotification(invitation.InviteUserId, invitation.GroupId, message)
}