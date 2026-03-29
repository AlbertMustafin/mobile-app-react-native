// types.ts
import { Route } from 'react-navigation';

// Define types for the application
type AuthRoute = Route<'Login' | 'Register'>;
type HomeRoute = Route<'Home'>;
type AppRoute = AuthRoute | HomeRoute;

// Define types for the application state
interface AppState {
  auth: AuthState;
}

interface AuthState {
  token: string | null;
  user: User | null;
}

interface User {
  id: number;
  name: string;
  email: string;
}

// Define types for the navigation props
interface AppNavigationProp {
  navigate: (routeName: string, params?: any) => void;
  goBack: () => void;
}