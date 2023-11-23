package urls

/*
provide implementation for the following functions in go
def aggregate_url():
    return('https://api.robinhood.com/options/aggregate_positions/')


def chains_url(symbol):
    return('https://api.robinhood.com/options/chains/{0}/'.format(id_for_chain(symbol)))


def option_historicals_url(id):
    return('https://api.robinhood.com/marketdata/options/historicals/{0}/'.format(id))


def option_instruments_url(id=None):
    if id:
        return('https://api.robinhood.com/options/instruments/{0}/'.format(id))
    else:
        return('https://api.robinhood.com/options/instruments/')


def option_orders_url(orderID=None, account_number=None):
    url = 'https://api.robinhood.com/options/orders/'
    if orderID:
        url += '{0}/'.format(orderID)
    if account_number:
        url += ('?account_numbers='+account_number)

    return url


def option_positions_url(account_number):
    if account_number:
        return('https://api.robinhood.com/options/positions/?account_numbers='+account_number)
    else:
        return('https://api.robinhood.com/options/positions/')


def marketdata_options_url():
    return('https://api.robinhood.com/marketdata/options/')

*/

// AggregateURL returns the URL for aggregate positions
func AggregateURL() string {
	return "https://api.robinhood.com/options/aggregate_positions/"
}
