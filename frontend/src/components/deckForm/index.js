import {useFormik} from 'formik';
import React, {useContext} from 'react';
import * as Yup from 'yup';
import {doUpload} from "../../service/api";
import {CurrentDeckContext, DeckListContext} from "../../context/currentDeck";

const FILE_SIZE = 60000000;

const SUPPORTED_FORMATS = [
  'application/pdf',
  'application/ppt',
];

function DeckForm(){
  const { decks, setDecks } = useContext(DeckListContext);
  const { setCurrentDeck } = useContext(CurrentDeckContext)

  const formik = useFormik({
    initialValues: {
      companyName: '',
      fileInput: null,
    },
    validationSchema: Yup.object({
      companyName: Yup.string().required('A company name is required'),
      fileInput: Yup.mixed()
        .required('A file is required')
        .test("fileSize", "The file is too large", (value) => {
          return value ? value.size <= FILE_SIZE : null
        })
        .test("fileType", "Only PDF's and PPT files are accepted", (value) => value ? SUPPORTED_FORMATS.includes(value.type) : null )
    }),
    onSubmit: async function onSubmit(values, {setSubmitting, resetForm}) {
      try {
        const resp = await doUpload(values);
        const d = decks
        d.push(resp)
        setDecks(d)
        setCurrentDeck(resp)
        resetForm()
      } catch(err) {
        for (let [key, value] of Object.entries(err)) {
          formik.errors[key] = value
        }
      }
      setSubmitting(false)
    }
  });

  return (
    <form onSubmit={formik.handleSubmit}>
      <div className="mb-3">
        <label htmlFor="nameInput" className="form-label">Company name</label>
        <input
          id="companyName"
          name="companyName"
          onChange={formik.handleChange}
          value={formik.values.companyName}
          type="text"
          className="form-control"
          aria-describedby="nameInputHelp"/>
          {formik.touched.companyName && formik.errors.companyName ? (
           <div className={"small text-danger"}>{formik.errors.companyName}</div>
         ) :  <div id="nameInputHelp" className="form-text">e.g. "Miraculous Melodies".</div> }
      </div>
      <div className="mb-3">
        <label htmlFor="fileInput" className="form-label">Your deck</label>
        <input
          id="fileInput"
          name="fileInput"
          onChange={(event) => {
            formik.setFieldValue("fileInput", event.currentTarget.files[0]);
          }}
          type="file"
          className="form-control"
          aria-describedby="fileInputHelp"/>
        {formik.touched.fileInput && formik.errors.fileInput ? (
          <div className={"small text-danger"}>{formik.errors.fileInput}</div>
        ) : <div id="fileInputHelp" className="form-text">Upload a PDF file to get started.</div> }
      </div>
      <button type="submit" className="btn btn-primary">Submit</button>
    </form>
  );
}

export default DeckForm