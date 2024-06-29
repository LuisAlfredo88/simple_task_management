import { validator } from "@/modules/shared/domain/errors";

export type User  = {
    id: string;
    userName: string;
    lastName: string;
    name: string;
    isActive: boolean;
    password: string;    
}

export type UserForm = User & {
    confirmationPassword: string;
    privileges: number[];
    roles: number[];
}

export type UserFilter = CriteriaFilter & {
    search?: string;
    status: number; 
}

export const validateUser = (user: UserForm) => {
    const dataValidation = [
        { validation: !user.name, error: 'SECURITY.must_specify_name', tag: 'name' },
        { validation: !user.lastName, error: 'SECURITY.must_specify_last_name', tag: 'last_name' },
        { validation: !user.userName, error: 'SECURITY.must_specify_user_name', tag: 'user_name' },
        { validation: !user.id && !user.password, error: 'SECURITY.must_specify_password', tag: 'password' },
        { 
            validation: !user.id && user.password != user.confirmationPassword, 
            error: 'SECURITY.password_confirmation_not_match' 
        },
    ];

    validator(dataValidation);
}
