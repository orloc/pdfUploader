import {useContext} from "react";
import {CurrentDeckContext} from "../../context/currentDeck";



function DeckList(){
  let decks = [
    {
      companyName: 'thing',
      uploadedAt: new Date(),
      numberOfSlides: 3,
      images: [
        "https://picsum.photos/200/200",
        "https://picsum.photos/200/200",
        "https://picsum.photos/200/200",
      ],
      id: 1
    },
    {
      companyName: 'stf',
      uploadedAt: new Date(),
      numberOfSlides: 6,
      images: [
        "https://picsum.photos/200/200",
        "https://picsum.photos/200/200",
        "https://picsum.photos/200/200",
        "https://picsum.photos/200/200",
      ],
      id: 2
    },
    {
      companyName: 'meows',
      uploadedAt: new Date(),
      numberOfSlides: 20,
      images: [
        "https://picsum.photos/200/200",
        "https://picsum.photos/200/200",
        "https://picsum.photos/200/200",
        "https://picsum.photos/200/200",
        "https://picsum.photos/200/200",
        "https://picsum.photos/200/200",
        "https://picsum.photos/200/200",
        "https://picsum.photos/200/200",
      ],
      id: 3
    },
  ];

  const { currentDeck, setCurrentDeck } = useContext(CurrentDeckContext)

  const selectDeck = (deck) => {
    setCurrentDeck(deck);
  };

  return (
    <ul className="list-group list-group-horizontal deckList">
      {decks.map((deck) => {
        return (
          <li key={`list-${deck.id}`} className={"list-group-item flex-fill " + (currentDeck && currentDeck.id === deck.id ? 'selected' : '')}
              onClick={() => selectDeck(deck)}
          >
            <div className="d-flex w-100 justify-content-between">
              <h6 className="mb-1">{deck.companyName}</h6>
              <small>{deck.uploadedAt.toLocaleDateString() }</small>
            </div>
            <small>{deck.numberOfSlides} images</small>
          </li>
        );
      })}
    </ul>
  );
}


export default DeckList