import {useContext} from "react";
import {CurrentDeckContext} from "../../context/currentDeck";

function DeckViewer(){
  const {currentDeck} = useContext(CurrentDeckContext)
  return (
    <div className="image-container mt-2">
      {currentDeck && currentDeck.images.map((img, i) => {
        return (<img key={`${i}-img`} src={img} className="rounded" />)
      })}
    </div>
  );
}

export default DeckViewer