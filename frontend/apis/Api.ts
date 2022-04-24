import axios from "axios"
import { CommonResponseType } from "./ApiType"

/**
 * 全API
 */
export class Api {
  // GET
  static hello = async (): Promise<CommonResponseType<string>> => {
    const res = await axios("localhost:8890/v1/hello")
    return res.data
  }


  // POST
}