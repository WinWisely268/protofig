package utils

import "fmt"

func isWhiteSpace(r rune) bool {
	return r == '\t' || r == ' '
}

func isNewLine(r rune) bool {
	return r == '\n' || r == '\r'
}

func isValidChar(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9') ||
		r == '_' ||
		r == '-'
}

func (ttyp tknType) String() string {
	switch ttyp {
	case tknEOF:
		return "EOF"
	case tknNIL:
		return "NIL"
	case tknError:
		return "Error"
	case tknString, tknMultilineString:
		return "String"
	case tknCommentStart:
		return "CommentStart"
	case tknKeyStart:
		return "KeyStart"
	case tknKeyEnd:
		return "KeyEnd"
	case tknCompStart:
		return "ComponentStart"
	case tknCompEnd:
		return "ComponentEnd"
	case tknText:
		return "Text"
	}
	panic(errUnknownTokenType)
}

func (t tkn) String() string {
	return fmt.Sprintf("%s, %s", t.typ.String(), t.val)
}
