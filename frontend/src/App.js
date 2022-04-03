import './App.css';
import DeckForm from "./components/deckForm";
import DeckViewer from "./components/deckViewer";
import DeckList from "./components/deckList";
import {useMemo, useState} from "react";
import {CurrentDeckContext, DeckListContext} from "./context/currentDeck";

function App() {
  const [currentDeck, setCurrentDeck] = useState(null);
  const [decks, setDecks] = useState(null);

  const value = useMemo(() => ({currentDeck, setCurrentDeck}), [currentDeck]);
  const deckValue = useMemo(() => ({decks, setDecks}), [decks]);

  return (
    <div className="App container">
      <CurrentDeckContext.Provider value={value}>
        <DeckListContext.Provider value={deckValue}>
        <div className="row">
          <div className="col-12">
            <div className="card mt-5">
              <div className="card-header">
                <h1>Pitch deck uploader 5000</h1>
              </div>
              <div className="card-body">
                <DeckForm/>
              </div>
            </div>
          </div>
        </div>
        <div className="row mt-2">
            <div className="col-12">
              <DeckList/>
            </div>
            <div className="col-12">
              <DeckViewer/>
            </div>
        </div>
      </DeckListContext.Provider>
      </CurrentDeckContext.Provider>
    </div>
  );
}

export default App;
