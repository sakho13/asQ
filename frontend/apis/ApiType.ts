export type CommonResponseType<T> = {
  result_flg: 1 | 0
  message: string
  response: T | null
}