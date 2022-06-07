import axios from "axios"
import { CommonResponseType } from "./ApiType"

/**
 * 全API
 */
export class Api {
  // GET
  /**
   * 接続確認用
   * @returns
   */
  static hello = async (): Promise<CommonResponseType<string>> => {
    const res = await axios.get("/v1/hello")
    return res.data
  }


  // POST
  static createUser = () => {
    const res = await axios.get("/v1/api/user/create")
    return res.data
  }
}