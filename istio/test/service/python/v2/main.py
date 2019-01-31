from apistar import App, Route, http
import platform
import requests
import logging


def getForwardHeaders(request):
    headers = {}
    incoming_headers = [
        'x-request-id',
        'x-b3-traceid',
        'x-b3-spanid',
        'x-b3-parentspanid',
        'x-b3-sampled',
        'x-b3-flags',
        'x-ot-span-context'
    ]

    for ihdr in incoming_headers:
        val = request.headers.get(ihdr)
        if val is not None:
            headers[ihdr] = val
            # print("incoming: "+ihdr+":"+val)

    return headers


# class CustomHeadersHook:
#     def on_response(self, request: http.Request, response: http.Response):
#         forwardHeaders = getForwardHeaders(request)
#         for k, v in forwardHeaders.items():
#             response.headers[k] = v
#         logging.debug(forwardHeaders)


def env(request: http.Request):
    forwardHeaders = getForwardHeaders(request)
    
    service_lua_url = 'http://' + 'service-lua' + '/env'
    resp = requests.get(service_lua_url, headers=forwardHeaders)
    data_lua = resp.json()

    service_node_url = 'http://' + 'service-node' + '/env'
    resp = requests.get(service_node_url, headers=forwardHeaders)
    data_node = resp.json()

    return {
        "message": 'Python' + platform.python_version() + '----->' + data_lua['message'] + ', ' + data_node['message']
    }


def status():
    return 'ok'


routes = [
    Route('/env', method='GET', handler=env),
    Route('/status', method='GET', handler=status),
]

# event_hooks = [CustomHeadersHook]
# app = App(routes=routes, event_hooks=event_hooks)
app = App(routes=routes)
logging.basicConfig(level=logging.DEBUG)

if __name__ == '__main__':
    app.serve('0.0.0.0', 80, debug=True)
