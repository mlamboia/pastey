package clipboard

import "github.com/atotto/clipboard"

func ReadClipboard() (string, error) {
    return clipboard.ReadAll()
}
