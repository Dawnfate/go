package main

type SubstituteRewardReq2 struct {
	Body `json:"BODY"`
}

type Body struct {
	EncryptData string `json:"ENCRYPT_DATA"`
	EncryptKey  string `json:"ENCRYPT_KEY"`
}

/*func main() {

		req := []byte("{\n    \"BODY\":{\n       \"ENCRYPT_DATA\":\"r8liyo0AFMZa2b2dZyJd6t9cBFN+b+x+aiy92jtqAV/EsEovaDrrC5m4ddqRBhJXj7YohKu6g8hAaqHKWeufuLb7bQbbTMIVMRUSzc4CcpiCA4P4LsZ5pBVHZAvlLIMO321yNHCndylQNGGnqOTNCVwBV1yw4YIElocn+kNNGIJyZyK7qBinMDGI14RNr8SqNc5W/COtm1K74T2hP8qGcLcTXyj1wrN2r6vOgwYXyrdArXzEv5I7L59nvEeKTMr7Qlm7Aka7cDcpoiRCLa+LhY/2uiwrWsV+LTT28+kLR10W9e2mLQ0217mmo/ICLk86wo/qrrrPWp4tHScjFjxboRiTOnJLxXJOANgsbJRQ6t9CztpZiWflooLvo2u5bSEM/OC0hr5revWI/Vo7OwGj2yBQoe7lJeYhAUisr3bEFIAI+7SmyMGYokMxyzRrmgvAHJNc4jyrj7YOZxv27KFTa2azk99m9Opp1inLYN1b9WnXnTe3zS58wgTfFH8lEROKh2251suh/Vj2BGNLShffpxGRZUWOlSe6C1mgeMexdyJ7rjiHkKotLz+lTa9TdMARWCTlKpPGol2H0gsZ4J+QhbSPmbkUGLB0wd7kjsICK/qO99DuhJwRcUW0VpXnaUtKXcPLmjph3DlJE+Q2+QZ+kFcxp7SAjDaE/7934pYfX9aE+DZr08clr+sv1FtOoZCXRE2P+hzoN2B8MD/0o8iYDRTRqsgUQUyFTSotsBZ25vmYOZjaM4//P4Se0vwV5GIOzjofj0fmd9Ofk1wKTnEU17w8YWDx3V1fiRsTPdogkw1qXl7ou/QI9Mhuemnwxd7OY4i/7U6KJb4rz6HZwAWbVYq2vmjRPu5/2dcNAG91HJv+U5iKiN8MZpybgbS35jwnN8ERT/GsdsnB1ZTfgz5eDB53hFILERtm37Ps7QROVV+6kjn6qg7Dp1rwAGPsh/La1W9DPDecGo+s73agqjpGPzpF8FiZ6/dJV+NDbak7QxL6wD+3T+dUsawH2RxzEfsF4EOi7Oky8PWTx6i6HDYjuw==\",\n       \"ENCRYPT_KEY\":\"I1Usy+4enLjkIZpABNSLFEmZRPfE/din9LbTklynw/eoHSCsEfYlF6hVVWEjdg1e1XOvlarSfUr5WvGMS5idqsYAPo2gz5FUeUk0Sv6i7q4iOvtWEwXRgxvH/OHkKFruBqR+EykvGeeQwg1ThkOPob9hZ/p2ITq5BfZzoGbyIUvDoXwzrmylp4UUQpPvV5wauDKwK44mHDHm2mqOYvA0N7SWryKEjEjHoeQshGxCfAqMM4PE4pMiGhHkZ/YLTvNaU0btz2SnkipAvsWhB2LAc5Ex1GTqGan/hXZyprSYvD5B2V4Tjrjtgh+59GVKdLyUFVBUlMeV99Ih26mvrT7z9Q==\"\n    }\n}")
		substituteRewardReq := new(SubstituteRewardReq2)
		err := json.Unmarshal(req, substituteRewardReq)
		if err != nil {
			panic(err)
		}
	aesKey := crypt.RSA_Decrypt([]byte("I1Usy+4enLjkIZpABNSLFEmZRPfE/din9LbTklynw/eoHSCsEfYlF6hVVWEjdg1e1XOvlarSfUr5WvGMS5idqsYAPo2gz5FUeUk0Sv6i7q4iOvtWEwXRgxvH/OHkKFruBqR+EykvGeeQwg1ThkOPob9hZ/p2ITq5BfZzoGbyIUvDoXwzrmylp4UUQpPvV5wauDKwK44mHDHm2mqOYvA0N7SWryKEjEjHoeQshGxCfAqMM4PE4pMiGhHkZ/YLTvNaU0btz2SnkipAvsWhB2LAc5Ex1GTqGan/hXZyprSYvD5B2V4Tjrjtgh+59GVKdLyUFVBUlMeV99Ih26mvrT7z9Q=="), "./private.pem")
	fmt.Println(string(aesKey))
	data := "r8liyo0AFMZa2b2dZyJd6t9cBFN+b+x+aiy92jtqAV/EsEovaDrrC5m4ddqRBhJXj7YohKu6g8hAaqHKWeufuLb7bQbbTMIVMRUSzc4CcpiCA4P4LsZ5pBVHZAvlLIMO321yNHCndylQNGGnqOTNCVwBV1yw4YIElocn+kNNGIJyZyK7qBinMDGI14RNr8SqNc5W/COtm1K74T2hP8qGcLcTXyj1wrN2r6vOgwYXyrdArXzEv5I7L59nvEeKTMr7Qlm7Aka7cDcpoiRCLa+LhY/2uiwrWsV+LTT28+kLR10W9e2mLQ0217mmo/ICLk86wo/qrrrPWp4tHScjFjxboRiTOnJLxXJOANgsbJRQ6t9CztpZiWflooLvo2u5bSEM/OC0hr5revWI/Vo7OwGj2yBQoe7lJeYhAUisr3bEFIAI+7SmyMGYokMxyzRrmgvAHJNc4jyrj7YOZxv27KFTa2azk99m9Opp1inLYN1b9WnXnTe3zS58wgTfFH8lEROKh2251suh/Vj2BGNLShffpxGRZUWOlSe6C1mgeMexdyJ7rjiHkKotLz+lTa9TdMARWCTlKpPGol2H0gsZ4J+QhbSPmbkUGLB0wd7kjsICK/qO99DuhJwRcUW0VpXnaUtKXcPLmjph3DlJE+Q2+QZ+kFcxp7SAjDaE/7934pYfX9aE+DZr08clr+sv1FtOoZCXRE2P+hzoN2B8MD/0o8iYDRTRqsgUQUyFTSotsBZ25vmYOZjaM4//P4Se0vwV5GIOzjofj0fmd9Ofk1wKTnEU17w8YWDx3V1fiRsTPdogkw1qXl7ou/QI9Mhuemnwxd7OY4i/7U6KJb4rz6HZwAWbVYq2vmjRPu5/2dcNAG91HJv+U5iKiN8MZpybgbS35jwnN8ERT/GsdsnB1ZTfgz5eDB53hFILERtm37Ps7QROVV+6kjn6qg7Dp1rwAGPsh/La1W9DPDecGo+s73agqjpGPzpF8FiZ6/dJV+NDbak7QxL6wD+3T+dUsawH2RxzEfsF4EOi7Oky8PWTx6i6HDYjuw=="
	subsData, err := crypt.AesDecrypt([]byte(data), aesKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(subsData)
}
*/
