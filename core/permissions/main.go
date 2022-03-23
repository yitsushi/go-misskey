package permissions

import "fmt"

// Permission is a string containing write or read permissions on a Resource.
type Permission string

// Resource represents a resource that can be accessed.
type Resource string

// Resource constants to make it easier to define permissions.
const (
	Account       Resource = "account"
	Blocks        Resource = "blocks"
	Channels      Resource = "channels"
	Drive         Resource = "drive"
	Favorites     Resource = "favorites"
	Following     Resource = "following"
	Gallery       Resource = "gallery"
	GalleryLikes  Resource = "gallery-likes"
	Messaging     Resource = "messaging"
	Mutes         Resource = "mutes"
	Notes         Resource = "notes"
	Notifications Resource = "notifications"
	PageLikes     Resource = "page-likes"
	Pages         Resource = "pages"
	Reactions     Resource = "reactions"
	UserGroups    Resource = "user-groups"
	Votes         Resource = "votes"
)

// Write permissions on given Resource.
func Write(area Resource) Permission {
	return Permission(fmt.Sprintf("write:%s", area))
}

// Read permissions on given Resource.
func Read(area Resource) Permission {
	return Permission(fmt.Sprintf("read:%s", area))
}
