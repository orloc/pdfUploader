const baseUrl = "http://localhost:3030"

function doUpload(data) {
  let fData = new FormData();
  fData.append("companyName", data.companyName)
  fData.append("file", data.fileInput)

  return fetch(`${baseUrl}/deck`, {
    method: "POST",
    body: fData,
  }).then((res) => {
    if (res.ok) {
      return { status: res.status }
    } else {
      return { status: res.status }
    }
  }).catch((err) => {
    return err;
  });
}

export {
  doUpload
}