import type { User } from './user'

export interface Auth {
    user: User;
    token: string;
    menu: Record<string, unknown>[],
}
