package helper

/*
Implement the following funtion using in go synthax, I have a global variable like sessions in globals/globals.go
L
def request_get(url, dataType='regular', payload=None, jsonify_data=True):
    """For a given url and payload, makes a get request and returns the data.

    :param url: The url to send a get request to.
    :type url: str
    :param dataType: Determines how to filter the data. 'regular' returns the unfiltered data. \
    'results' will return data['results']. 'pagination' will return data['results'] and append it with any \
    data that is in data['next']. 'indexzero' will return data['results'][0].
    :type dataType: Optional[str]
    :param payload: Dictionary of parameters to pass to the url. Will append the requests url as url/?key1=value1&key2=value2.
    :type payload: Optional[dict]
    :param jsonify_data: If this is true, will return requests.post().json(), otherwise will return response from requests.post().
    :type jsonify_data: bool
    :returns: Returns the data from the get request. If jsonify_data=True and requests returns an http code other than <200> \
    then either '[None]' or 'None' will be returned based on what the dataType parameter was set as.

    """
    if (dataType == 'results' or dataType == 'pagination'):
        data = [None]
    else:
        data = None
    res = None
    if jsonify_data:
        try:
            res = SESSION.get(url, params=payload)
            res.raise_for_status()
            data = res.json()
        except (requests.exceptions.HTTPError, AttributeError) as message:
            print(message, file=get_output())
            return(data)
    else:
        res = SESSION.get(url, params=payload)
        return(res)
    # Only continue to filter data if jsonify_data=True, and Session.get returned status code <200>.
    if (dataType == 'results'):
        try:
            data = data['results']
        except KeyError as message:
            print("{0} is not a key in the dictionary".format(message), file=get_output())
            return([None])
    elif (dataType == 'pagination'):
        counter = 2
        nextData = data
        try:
            data = data['results']
        except KeyError as message:
            print("{0} is not a key in the dictionary".format(message), file=get_output())
            return([None])

        if nextData['next']:
            print('Found Additional pages.', file=get_output())
        while nextData['next']:
            try:
                res = SESSION.get(nextData['next'])
                res.raise_for_status()
                nextData = res.json()
            except:
                print('Additional pages exist but could not be loaded.', file=get_output())
                return(data)
            print('Loading page '+str(counter)+' ...', file=get_output())
            counter += 1
            for item in nextData['results']:
                data.append(item)
    elif (dataType == 'indexzero'):
        try:
            data = data['results'][0]
        except KeyError as message:
            print("{0} is not a key in the dictionary".format(message), file=get_output())
            return(None)
        except IndexError as message:
            return(None)

    return(data)
*/

// RequestGet is a function that makes a get request and returns the data and error if any.
func RequestGet(url string, dataType string, payload map[string]string, jsonifyData bool) (interface{}, error) {
	//use the global variable Session to make the request and implement the logic of the function
	net.
}
