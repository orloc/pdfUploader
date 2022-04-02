import './App.css';
import DeckForm from "./components/deckForm";
import DeckViewer from "./components/deckViewer";
import DeckList from "./components/deckList";

function App() {
  return (
    <div className="App" className="container">
      <div className="row">
        <div className="col-12">
          <div className="card mt-5">
            <div className="card-header">
                <h1>Pitch deck uploader 5000</h1>
            </div>
            <div className="card-body">
              { DeckForm() }
            </div>
          </div>
        </div>
      </div>
      <div className="row mt-2">
        <div className="col-12">
          { DeckList() }
        </div>
        <div className="col-12">
          { DeckViewer() }
        </div>
      </div>
    </div>
  );
}

export default App;
