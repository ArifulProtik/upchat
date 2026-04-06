export interface UserResponse {
	id: string;
	created_at: string;
	updated_at: string;
	name: string;
	email: string;
}

export interface AccountResponse {
	provider: string;
	mail_verified: boolean;
}

export interface LoginResponse {
	token: string;
	user: UserResponse;
}

export interface GetSessionResponse {
	user: UserResponse;
	account: AccountResponse;
}

export interface FieldError {
	field: string;
	message: string;
}

export interface ApiErrorBody {
	status: number;
	error: string | { errors: FieldError[] };
}
