import {useContext, useEffect, useState} from "react";
import {CurrentDeckContext, DeckListContext} from "../../context/currentDeck";
import {getDecks} from "../../service/api";

function DeckList(){
  const { decks, setDecks } = useContext(DeckListContext);
  const { currentDeck, setCurrentDeck } = useContext(CurrentDeckContext)

  const selectDeck = (deck) => {
    setCurrentDeck(deck);
  };

  useEffect(() => {
    const fetchData = async () => {
      const data = await getDecks();
      return data;
    };
    fetchData().then((res) => {
      setDecks(res);
    }).catch((err) => {
      console.log(err);
    });
  }, [])

  return (
    <ul className="list-group list-group-horizontal deckList">
      {decks && decks.map((deck) => {
        return (
          <li key={`list-${deck.uuid}`} className={"list-group-item flex-fill " + (currentDeck && currentDeck.uuid === deck.uuid ? 'selected' : '')}
              onClick={() => selectDeck(deck)}
          >
            <div className="d-flex w-100 justify-content-between">
              <h5 className="mb-1">{deck.companyName}</h5>
            </div>
            <small>Uploaded at: {deck.createdAt.toLocaleDateString() }</small>
            <small>{deck.numberOfSlides} images</small>
          </li>
        );
      })}
    </ul>
  );
}

export default DeckList