// var myRemote = 'http://localhost:4006'
var myRemote = ''

export function getRecipes (id) {
  return new Promise((resolve, reject) => {
    var xhr = new XMLHttpRequest()
    xhr.onload = () => {
      console.log('onload', xhr.status, xhr.responseType, xhr.response)
      if (xhr.status !== 200 /* ok */) {
        reject(new Error(xhr.statusText + ' : ' + xhr.responseText))
      } else {
        try {
          var recipes = JSON.parse(xhr.responseText)
          resolve(recipes)
        } catch (err) {
          reject(err)
        }
      }
    }
    xhr.onerror = () => reject(new Error('an error occurred during the transaction'))
    xhr.ontimeout = (e) => reject(new Error('timeout'))
    var optionalID = (typeof id === 'string' && id !== '') ? '/' + id : ''
    xhr.open('GET', myRemote + '/api/recipes' + optionalID, true)
    xhr.send()
  })
}

export function deleteRecipe (id) {
  return new Promise((resolve, reject) => {
    var xhr = new XMLHttpRequest()
    xhr.onload = () => {
      console.log('onload', xhr.status, xhr.statusText, xhr.responseType, xhr.response)
      if (xhr.status === 204 /* No Content */) {
        console.log('ok', xhr.response)
        resolve(xhr.statusText)
      } else {
        reject(new Error(xhr.statusText + ' : ' + xhr.responseText))
      }
    }
    xhr.onerror = () => reject(new Error('an error occurred during the transaction'))
    xhr.ontimeout = (e) => reject(new Error('timeout'))
    xhr.open('DELETE', myRemote + '/api/recipes/' + id, true)
    xhr.send()
  })
}

export function postOrPutRecipe (method, recipe) {
  console.log('postOrPutRecipe')
  return new Promise((resolve, reject) => {
    var xhr = new XMLHttpRequest()
    xhr.onload = () => {
      console.log('onload', xhr.status, xhr.statusText, xhr.responseType, xhr.response)
      if (xhr.status === 201 /* created */) {
        var location = xhr.getResponseHeader('Location')
        resolve(location)
      } else {
        reject(new Error(xhr.statusText + ' : ' + xhr.responseText))
      }
    }
    xhr.onerror = () => reject(new Error('an error occurred during the transaction'))
    xhr.ontimeout = (e) => reject(new Error('timeout'))
    // if this is an update of existing recipe, we need to PUT.
    // in that case recipe_id is expected to be set and should be included in URL
    var idStr = ''
    if (typeof recipe.recipe_id === 'string' && recipe.recipe_id !== '' && method === 'PUT') {
      idStr = '/' + idStr
    }
    xhr.open(method, myRemote + '/api/recipes' + idStr, true)
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.send(JSON.stringify(recipe))
  })
}
