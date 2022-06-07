/**
 * 全API共通レスポンス型
 */
export type CommonResponseType<T> = {
  // 1: success, 0: err
  result_flg: 1 | 0
  message: string
  response: T | null
}

/**
 * /v1/api/user/create
 */
export type CreateUserInput = {
  firebase_jwt: string
}
