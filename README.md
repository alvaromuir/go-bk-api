## go-bk-api
### A Golang Library for the Oracle / BlueKai API


Requires the following environmental variables:

* BK_KEY
* BK_SECRET
* BK_PARTNER_ID

The BK_KEY and BK_SECRET refer to BlueKai's Web Service User Key, specific to your login. Please refer to Oracle's [documentation](http://docs.oracle.com/cloud/latest/marketingcs_gs/OMCDB/#Developers/api_reference.html) to learn how to obtain these.


Requires:
* [GoDotEnv](https://github.com/joho/godotenv)


For convenience, these can be set in a `.env` file placed in the project root. An example live in `.env-template`

A `examples.go` file with a few test functions have been included for illustrative purposes.

Alvaro Muir, [@alvaromuir](https://twitter.com/alvaromuir)
