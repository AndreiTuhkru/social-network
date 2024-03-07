// hooks/useAuthCheck.js

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useAuth } from '@/hooks/utils';

const useAuthCheck = () => {
    const router = useRouter();
    const checkAuth = async () => {
        const auth = await useAuth();
        if (!auth.is_authenticated) {
            router.push('/auth');
        }
    };

    useEffect(() => {
        checkAuth();
        // Optionally set up a recurring check
        const intervalId = setInterval(checkAuth, 1800000); // 30 minutes
        return () => clearInterval(intervalId);
    }, [router]);
};

export default useAuthCheck;