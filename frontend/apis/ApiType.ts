/**
 * 全API共通レスポンス型
 */
export type CommonResponseType<T> = {
  result_flg: 1 | 0
  message: string
  response: T | null
}