/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

type security struct {
	xpBlackList     map[[16]byte]*blackList
	routerBlackList map[[16]byte]*blackList
	userBlackList   map[[16]byte]*blackList
}

type blackList struct {

}