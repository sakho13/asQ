import axios from "axios";
import {
  CommonResponseType,
  CreateUserInput,
  CreateUserOutput,
} from "./ApiType";

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
    const res = await axios.get("/v1/hello");
    return res.data;
  };

  // POST

  static createUser = async (
    input: CreateUserInput
  ): Promise<CreateUserOutput> => {
    const res = await axios.post("/v1/api/user/create", input);
    return res.data as CreateUserOutput;
  };
}
