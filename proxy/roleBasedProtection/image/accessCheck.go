package image

import "fmt"

type AccessCheck struct {
	filename string
	userRole string
	image    *Image
}

func NewAccessCheck(filename, userRole string) AccessCheck {
	return AccessCheck{
		filename: filename,
		userRole: userRole,
	}
}

func (this AccessCheck) GetFileName() string {
	return this.filename
}

func (this AccessCheck) Display() {
	if !this.checkAccess() {
		fmt.Println("SecureImageProxy: ACCESS DENIED for " + this.filename + " (role: " + this.userRole + ")")
		return
	}
	if this.image == nil {
		this.image = NewImage(this.filename)
	}
	this.image.Display()

}

func (this AccessCheck) checkAccess() bool {
	fmt.Println("SecureImageProxy: Checking access for role '" + this.userRole + "' on " + this.filename)
	return this.userRole == "ADMIN"
}
