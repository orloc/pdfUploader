const baseUrl = "http://localhost:3030"

async function doUpload(data) {
  let fData = new FormData();
  fData.append("companyName", data.companyName)
  fData.append("file", data.fileInput)

  const response = await fetch(`${baseUrl}/deck`, {
    method: "POST",
    body: fData,
    headers: {
      'Accept': 'application/json'
    }
  });

  if (response.status === 409) {
    throw { companyName: "Company already exists." }
  }

  const resp = await response.text()
  let js = JSON.parse(resp)
  return parseDeckResponse(js, true);
}

async function getDecks() {
  const response = await fetch(`${baseUrl}/decks`, {
    method: "GET",
    headers: {
      'Accept': 'application/json'
    }
  });

  const resp = await response.text()
  return parseDeckResponse(JSON.parse(resp));
}

function parseDeckResponse(resp, single) {
  if (!resp || !resp.data) return single ? null : [];
  if (!single) {
    return resp.data.map((d) => {
      return transformDesk(d)
    })
  }
  return transformDesk(resp.data)
}

function transformDesk(d) {
  let { company_name, uuid, created_at, images} = d

  images = images.map((img) => `http://localhost:3030/${img}`)

  return {
    uuid,
    images,
    numberOfSlides: images ? images.length : 0,
    companyName : company_name,
    createdAt: new Date(created_at),
  }
}

export {
  doUpload,
  getDecks
}