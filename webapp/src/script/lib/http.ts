type JSON = any

function makeRequest(method: string, route: string, data: string, contentType: string): Promise<XMLHttpRequestResponseType> {
    return new Promise(function (resolve, reject) {
        const url = window.location.origin + route
        console.log(url)
        let xhr = new XMLHttpRequest();
        xhr.open(method, url);
        xhr.setRequestHeader("Content-Type", contentType);
        xhr.onload = function () {
            if (this.status >= 200 && this.status < 300) {
                resolve(xhr.response);
            } else {
                reject({
                    status: this.status,
                    statusText: xhr.statusText
                });
            }
        };
        xhr.onerror = function () {
            reject({
                status: this.status,
                statusText: xhr.statusText
            });
        };
        xhr.send(data);
    });
}

async function jsonRequestJsonResponse(method: string, route: string, json: JSON): Promise<JSON | null> {
    const resp = await makeRequest(method, route, JSON.stringify(json), "application/json;charset=UTF-8")
    return JSON.parse(resp)
}

export { jsonRequestJsonResponse }
