/* For license and copyright information please see the LEGAL file in the code repository */

package os_p

// TODO::: Why we need signals when OS can call apps services that force by OS to be in the app??

// Signal_Listener
type Signal_Listener interface {
	// Non-Blocking, means It must not block the caller in any ways.
	// https://en.wikipedia.org/wiki/Signal_(IPC)
	OsSignalHandler(signal Signal)
}

type Signal int

const (
	Signal_Unset Signal = iota

	// Signal_Hangup re-purpose as a signal to re-read configuration files, or reinitialize
	// https://en.wikipedia.org/wiki/SIGHUP
	Signal_Hangup // syscall.Signal(10): // syscall.SIGUSR1

	Signal_Interrupt
	Signal_Quit
	Signal_IllegalInstruction
	Signal_TraceTrap // "trace/breakpoint trap"
	Signal_Aborted
	Signal_BusError
	Signal_FloatingPointException
	Signal_Killed
	Signal_SegmentationFault
	Signal_BrokenPipe
	Signal_AlarmClock
	Signal_Terminated
	Signal_Urgent // syscall.Signal(0x17)

	Signal_Upgrade

	// "user defined signal 1",
	// "user defined signal 2",

	// https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/signal.h.html
	/*
			SIGABRT		A		Process abort signal.
			SIGALRM		T		Alarm clock.
			SIGBUS		A		Access to an undefined portion of a memory object.
			SIGCHLD		I		Child process terminated, stopped,
			[XSI] [Option Start]		or continued. [Option End]
			SIGCONT		C		Continue executing, if stopped.
			SIGFPE		A		Erroneous arithmetic operation.
			SIGHUP		T		Hangup.
			SIGILL		A		Illegal instruction.
			SIGINT		T		Terminal interrupt signal.
			SIGKILL		T		Kill (cannot be caught or ignored).
			SIGPIPE		T		Write on a pipe with no one to read it.
			SIGQUIT		A		Terminal quit signal.
			SIGSEGV		A		Invalid memory reference.
			SIGSTOP		S		Stop executing (cannot be caught or ignored).
			SIGTERM		T		Termination signal.
			SIGTSTP		S		Terminal stop signal.
			SIGTTIN		S		Background process attempting read.
			SIGTTOU		S		Background process attempting write.
			SIGUSR1		T		User-defined signal 1.
			SIGUSR2		T		User-defined signal 2.
			[OB XSR] [Option Start]
			SIGPOLL		T		Pollable event. [Option End]
			[OB XSI] [Option Start]
			SIGPROF		T		Profiling timer expired. [Option End]
			[XSI] [Option Start]
			SIGSYS		A		Bad system call. [Option End]
			SIGTRAP		A		Trace/breakpoint trap. [Option End]
			SIGURG		I		High bandwidth data is available at a socket.
			[XSI] [Option Start]
			SIGVTALRM	T		Virtual timer expired.
			SIGXCPU		A		CPU time limit exceeded.
			SIGXFSZ		A		File size limit exceeded. [Option End]

		T	Abnormal termination of the process.
		A	Abnormal termination of the process [XSI] [Option Start]  with additional actions. [Option End]
		I	Ignore the signal.
		S	Stop the process.
		C	Continue the process, if it is stopped; otherwise, ignore the signal.
	*/

	/*
			https://man7.org/linux/man-pages/man7/signal.7.html

		       SIGABRT      P1990      Core    Abort signal from abort(3)
		       SIGALRM      P1990      Term    Timer signal from alarm(2)
		       SIGBUS       P2001      Core    Bus error (bad memory access)
		       SIGCHLD      P1990      Ign     Child stopped or terminated
		       SIGCLD         -        Ign     A synonym for SIGCHLD
		       SIGCONT      P1990      Cont    Continue if stopped
		       SIGEMT         -        Term    Emulator trap
		       SIGFPE       P1990      Core    Floating-point exception
		       SIGHUP       P1990      Term    Hangup detected on controlling terminal
		                                       or death of controlling process
		       SIGILL       P1990      Core    Illegal Instruction
		       SIGINFO        -                A synonym for SIGPWR
		       SIGINT       P1990      Term    Interrupt from keyboard

		       SIGIO          -        Term    I/O now possible (4.2BSD)
		       SIGIOT         -        Core    IOT trap. A synonym for SIGABRT
		       SIGKILL      P1990      Term    Kill signal
		       SIGLOST        -        Term    File lock lost (unused)
		       SIGPIPE      P1990      Term    Broken pipe: write to pipe with no
		                                       readers; see pipe(7)
		       SIGPOLL      P2001      Term    Pollable event (Sys V);
		                                       synonym for SIGIO
		       SIGPROF      P2001      Term    Profiling timer expired
		       SIGPWR         -        Term    Power failure (System V)
		       SIGQUIT      P1990      Core    Quit from keyboard
		       SIGSEGV      P1990      Core    Invalid memory reference
		       SIGSTKFLT      -        Term    Stack fault on coprocessor (unused)
		       SIGSTOP      P1990      Stop    Stop process
		       SIGTSTP      P1990      Stop    Stop typed at terminal
		       SIGSYS       P2001      Core    Bad system call (SVr4);
		                                       see also seccomp(2)
		       SIGTERM      P1990      Term    Termination signal
		       SIGTRAP      P2001      Core    Trace/breakpoint trap
		       SIGTTIN      P1990      Stop    Terminal input for background process
		       SIGTTOU      P1990      Stop    Terminal output for background process
		       SIGUNUSED      -        Core    Synonymous with SIGSYS
		       SIGURG       P2001      Ign     Urgent condition on socket (4.2BSD)
		       SIGUSR1      P1990      Term    User-defined signal 1
		       SIGUSR2      P1990      Term    User-defined signal 2
		       SIGVTALRM    P2001      Term    Virtual alarm clock (4.2BSD)
		       SIGXCPU      P2001      Core    CPU time limit exceeded (4.2BSD);
		                                       see setrlimit(2)
		       SIGXFSZ      P2001      Core    File size limit exceeded (4.2BSD);
		                                       see setrlimit(2)
		       SIGWINCH       -        Ign     Window resize signal (4.3BSD, Sun)
	*/
)
