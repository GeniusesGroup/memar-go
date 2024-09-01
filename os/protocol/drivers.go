/* For license and copyright information please see the LEGAL file in the code repository */

package os_p

type Driver interface {
	Open()  // shall turn the hardware on		Enable()	||
	Close() // shall turn the hardware off		Disable()	||
	// init() // allocate and init the device
	// deinit() // cleanup
}

// https://www.kernel.org/doc/Documentation/input/input.txt
type InputDriver interface {
	Time() int64
	Type() uint8 // event type e.g. EV_REL for relative moment, EV_KEY for a keypress or release	More types are defined in include/uapi/linux/input-event-codes.h
	Code() uint8 // event code e.g. REL_X, KEY_BACKSPACE, ...										a complete list is in include/uapi/linux/input-event-codes.h.
	Value() int
}

type OutputDriver interface {
}

// NVM Express (NVMe) or Non-Volatile Memory Host Controller Interface
// https://en.wikipedia.org/wiki/SCSI_command
// https://spdk.io/doc/nvme.html
// https://github.com/linux-nvme/nvme-cli
type NVME interface {
	StorageBlockDetails
	Read()
	Erase()
	Write()
	Flush() // Flush queues to storage usually to remove device from computer
}

type StorageBlockDetails interface {
	RAID()
	VolumeID() UUID       // ID on OS level or platform level
	VolumeName()          // can be device name on OS level
	VolumeSize()          // in MB ??
	VolumeType()          // SSD, HDD, OldGeneration, ...
	VolumeMaxIOPS()       // 256000, 250, 40, ...
	VolumeMaxThroughput() // 4000, 250, 40, ...
	PhysicalSectorSize()  // 512, 4096
}
