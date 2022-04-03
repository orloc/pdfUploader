import {useCallback, useContext, useState} from "react";
import {CurrentDeckContext} from "../../context/currentDeck";
import ImageViewer from 'react-simple-image-viewer';

function DeckViewer(){
  const {currentDeck} = useContext(CurrentDeckContext)
  const [currentImage, setCurrentImage] = useState(0);
  const [isViewerOpen, setIsViewerOpen] = useState(false);

  const openImageViewer = useCallback((index) => {
    setCurrentImage(index);
    setIsViewerOpen(true);
  }, []);

  const closeImageViewer = () => {
    setCurrentImage(0);
    setIsViewerOpen(false);
  };

  return (
    <div className="image-container mt-2">
      {currentDeck && currentDeck.images.map((img, i) => {
        return (<img
          onClick={ () => openImageViewer((i))}
          key={`${i}-img`}
          src={img}
          className="rounded" />)
      })}

      { isViewerOpen && (
        <ImageViewer
          src={ currentDeck.images }
          currentIndex={ currentImage }
          disableScroll={ false }
          closeOnClickOutside={ true }
          onClose={ closeImageViewer }
        />
      )}
    </div>
  );
}

export default DeckViewer