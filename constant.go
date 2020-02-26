package pay_center_client

const (
	ACCOUNT_STATUS_NORMAL = 0 //账户正常
	ACCOUNT_STATUS_LOCKED = 1 //账户锁定

	ACCOUNT_TYPE_COMMON = 0 //通用预存款账户

	URL_RECHARGE_CASH_PAY         = "http://127.0.0.1:8128/pay-center/recharge/cash/pay"
	URL_RECHARGE_WEIXIN_GOPAY     = "http://127.0.0.1:8128/pay-center/recharge/weixin/gopay"
	URL_BILL_DEDUCT_AUTO          = "http://127.0.0.1:8128/pay-center/bill/deduct/auto"
	URL_BILL_CASH_PAY             = "http://127.0.0.1:8128/pay-center/bill/cash/pay"
	URL_BILL_WEIXIN_GOPAY         = "http://127.0.0.1:8128/pay-center/bill/weixin/gopay"
	URL_ACCOUNT_CREATE            = "http://127.0.0.1:8128/pay-center/account/add"
	URL_ACCOUNT_UPDATE            = "http://127.0.0.1:8128/pay-center/account/update"
	URL_ACCOUNT_GET               = "http://127.0.0.1:8128/pay-center/account/get"
	URL_ACCOUNT_FLOW_LIST         = "http://127.0.0.1:8128/pay-center/account/flow/list"
	URL_REFUND_TO_ACCOUNT         = "http://127.0.0.1:8128/pay-center/refund/to/account"
	URL_REFUND_TO_USER            = "http://127.0.0.1:8128/pay-center/refund/to/user"
	URL_STATISTICS_RECHARGE_DAILY = "http://127.0.0.1:8128/pay-center/statistics/recharge/daily"
	URL_STATISTICS_RECHARGE_MONTH = "http://127.0.0.1:8128/pay-center/statistics/recharge/month"

	APPID  = "test"
	APIKEY = "testtesttesttesttesttesttesttest"
)
