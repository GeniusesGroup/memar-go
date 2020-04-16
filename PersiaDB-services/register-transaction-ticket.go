/* For license and copyright information please see LEGAL file in repository */

package services

// or transactional authority
// or index lock ticket!

// transaction write can be on secondary indexes not primary indexes, due to primary index must always unique!
// transaction manager on any node in a replication must sync with master replication corresponding node manager!

// Get a record by ID when record ready to submit! Usually use in transaction queue to act when record ready to read!
// Must send this request to specific node that handle that range!!