package domain

const defaultFileName = "data"
const defaultUrlFile = ""
const defaultSeparator = ","

type ReaderOption struct {
	fileName  string
	urlFile   string
	separator string
}

func NewReaderOption() *ReaderOption {
	return &ReaderOption{
		fileName:  defaultFileName,
		urlFile:   defaultUrlFile,
		separator: defaultSeparator,
	}
}
func (r *ReaderOption) SetFileName(fileName string) {
	r.fileName = fileName
}
func (r *ReaderOption) SetUrlFile(urlFile string) {
	r.urlFile = urlFile
}
func (r *ReaderOption) SetSeparator(separator string) {
	r.separator = separator
}
func (r *ReaderOption) GetFileName() string {
	return r.fileName
}
func (r *ReaderOption) GetUrlFile() string {
	return r.urlFile
}
func (r *ReaderOption) GetSeparator() string {
	return r.separator
}
