var myRemote = 'http://localhost:4006'

export function getRecipes () {
  return new Promise((resolve, reject) => {
    var xhr = new XMLHttpRequest()
    xhr.responseType = 'json'
    xhr.onload = () => {
      console.log('onload', xhr.status, xhr.statusText, xhr.responseType, xhr.response)
      if (xhr.response) {
        console.log('ok', xhr.response)
        resolve(xhr.response)
      } else {
        reject(new Error('bad response format ' +
          xhr.responseType))
      }
    }
    xhr.onerror = () => reject(new Error('an error occurred during the transaction'))
    xhr.ontimeout = (e) => reject(new Error('timeout'))
    xhr.open('GET', myRemote + '/api/recipes', true)
    xhr.send()
  })
}
