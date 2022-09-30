package music

type Playlist struct {
	Title string
	Songs []Song
}

func (playlist *Playlist) AddSong(song Song) {
	playlist.Songs = append(playlist.Songs, song)
}

func (playlist *Playlist) Length() int {
	if playlist.Songs == nil {
		return 0
	}
	return len(playlist.Songs)
}
