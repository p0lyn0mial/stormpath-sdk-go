package stormpath

const (
	GroupMembershipBaseUrl = "https://api.stormpath.com/v1/groupMemberships"
)

type GroupMembership struct {
	Href    string `json:"href,omitempty"`
	Account Link   `json:"account"`
	Group   Link   `json:group`
}

func NewGroupMembership(accountHref string, groupHref string) *GroupMembership {
	return &GroupMembership{Account: Link{accountHref}, Group: Link{groupHref}}
}
