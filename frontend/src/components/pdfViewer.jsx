import { Document, Page } from "react-pdf/dist/esm/entry.webpack";
import { useState } from "react";
import "../css/pdfViewer.css";

function PDFViewer({ pdf, thumbnail, onClick }) {
    const [numPages, setNumPages] = useState(null);

    function onDocumentLoadSuccess({ newNumPages }) {
        setNumPages(newNumPages);
    }

    return (
        <div
            className={`${thumbnail ? "thumbnail-wrapper" : ""} pdf-wrapper`}
            onClick={onClick}
        >
            <Document
                file={"data:application/pdf;base64," + pdf}
                onLoadSuccess={onDocumentLoadSuccess}
                className={thumbnail ? "" : "pdf-container"}
            >
                <Page
                    pageNumber={1}
                    className="pdf-page"
                    scale={thumbnail ? 0.6 : 1.5}
                />
            </Document>
        </div>
    );
}

export default PDFViewer;