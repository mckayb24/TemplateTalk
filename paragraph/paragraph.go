package paragraph

func GetName() string {
	return "Paragraph"
}

func GetFile() string {
	return "paragraph/paragraph.html"
}

func GetContent() []string {
	return []string{
		"	This is the first line.\n",
		"The first line was indented.\n",
		"The previous line wasn't indented, neither is this one.\n",
		"Unless this is the first time in which case it is all messed up.\n",
	}
}
