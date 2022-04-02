

function DeckForm(){
  return (
    <form>
      <div className="mb-3">
        <label htmlFor="nameInput" className="form-label">Company name</label>
        <input type="text" className="form-control" id="nameInput" aria-describedby="nameInputHelp"/>
        <div id="nameInputHelp" className="form-text">e.g. "Miraculous Melodies".</div>
      </div>
      <div className="mb-3">
        <label htmlFor="fileInput" className="form-label">Your deck</label>
        <input type="file" className="form-control" id="fileInput" aria-describedby="fileInputHelp"/>
        <div id="fileInputHelp" className="form-text">Upload a PDF file to get started.</div>
      </div>
      <button type="submit" className="btn btn-primary">Submit</button>
    </form>
  );

}

export default DeckForm