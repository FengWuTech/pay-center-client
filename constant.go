package pay_center_client

const (
	ACCOUNT_STATUS_NORMAL = 0 //账户正常
	ACCOUNT_STATUS_LOCKED = 1 //账户锁定

	ACCOUNT_TYPE_COMMON = 0 //通用预存款账户

	URL_RECHARGE_CASH_PAY     = "http://pms-api.gmtech.top/pay-center/minip-user/recharge/cash/pay"
	URL_RECHARGE_WEIXIN_GOPAY = "http://pms-api.gmtech.top/pay-center/minip-user/recharge/weixin/gopay"
	URL_BILL_DEDUCT_AUTO      = "http://pms-api.gmtech.top/pay-center/minip-user/bill/deduct/auto"
	URL_BILL_CASH_PAY         = "http://pms-api.gmtech.top/pay-center/minip-user/bill/cash/pay"
	URL_BILL_WEIXIN_GOPAY     = "http://pms-api.gmtech.top/pay-center/minip-user/bill/weixin/gopay"
	URL_ACCOUNT_CREATE        = "http://pms-api.gmtech.top/pay-center/minip-user/account/add"
	URL_ACCOUNT_UPDATE        = "http://pms-api.gmtech.top/pay-center/minip-user/account/update"
	URL_ACCOUNT_GET           = "http://pms-api.gmtech.top/pay-center/minip-user/account/get"
	URL_REFUND_TO_ACCOUNT     = "http://pms-api.gmtech.top/pay-center/minip-user/refund/to/account"
	URL_REFUND_TO_USER        = "http://pms-api.gmtech.top/pay-center/minip-user/refund/to/user"

	APPID  = "test"
	APIKEY = "testtesttesttesttesttesttesttest"
)
