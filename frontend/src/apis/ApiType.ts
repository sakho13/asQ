/**
 * 全API共通レスポンス型
 */
export type CommonResponseType<T> =
  | {
      // 1: success, 0: err
      result_flg: 1;
      message: string;
      response: T;
    }
  | {
      // 1: success, 0: err
      result_flg: 0;
      message: string;
      response: null;
    };

/**
 * INPUT
 *
 * /v1/api/user/create
 */
export type CreateUserInput = {
  firebase_jwt: string;
};

/**
 * OUTPUT
 *
 * /v1/api/user/create
 */
export type CreateUserOutput = CommonResponseType<{
  user_id: string;
}>;
