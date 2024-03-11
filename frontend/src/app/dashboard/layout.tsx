"use client"
import MainHeader from '@/components/headers/MainHeader';
import background from '../../../public/assets/background.png';
import Footer from '@/components/headers/Footer';
import { TextColorModeProvider } from '@/hooks/TextColorModeContext';

import {ChatProvider, ChatManager} from '@/components/chat/ChatContext';





export default function DashboardLayout({
    children,
  }: {
    children: React.ReactNode
  }) {
    return (
      <TextColorModeProvider>
        <ChatProvider>
        <div style={{
          backgroundImage: `url("${background.src}")`,
          backgroundSize: 'cover',
          flex: 1,

      }}>
        
        <MainHeader/>
        {children}
        <ChatManager/>
        <Footer/>
        </div>
      </ChatProvider>
      </TextColorModeProvider>
        
    )
  }

