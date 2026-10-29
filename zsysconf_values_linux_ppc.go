// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && ppc

// POSIX limits for 32-bit big-endian PowerPC Linux.
//
// NOT cgo -godefs output: the generator needs a C toolchain for the target, so
// these are transcribed from the 32-bit ILP32 sibling (linux/386), which agrees
// with the ppc64 file on all 97 constants except the ones that are genuinely
// word-size dependent (_POSIX_V6/V7_ILP32_* / _LP64_*, _PTHREAD_STACK_MIN), and
// those take the ILP32 values here.
//
// Provenance caveat, deliberate: these are glibc's values. A musl-based ppc32
// Linux (Adelie, this port's test venue) reports some of them differently. That
// is measured, not assumed: a probe run in the Adelie rootfs gives PAGESIZE 4096,
// LOGIN_NAME_MAX 256 (both agreeing with these values) but HOST_NAME_MAX 255,
// where glibc documents 64 - so SC_HOST_NAME_MAX answers 64 here where the
// target's own libc says 255. Nothing in syncthing's closure depends on that
// difference; fix the affected values from the target libc if a consumer ever
// does. PTHREAD_STACK_MIN is also libc-specific and was not probed.
package sysconf

const (
	_AIO_PRIO_DELTA_MAX = 0x14
	_BC_BASE_MAX        = 0x63
	_BC_DIM_MAX         = 0x800
	_BC_SCALE_MAX       = 0x63
	_BC_STRING_MAX      = 0x3e8
	_COLL_WEIGHTS_MAX   = 0xff
	_DELAYTIMER_MAX     = 0x7fffffff
	_EXPR_NEST_MAX      = 0x20
	_HOST_NAME_MAX      = 0x40
	_LINE_MAX           = 0x800
	_LOGIN_NAME_MAX     = 0x100
	_MQ_PRIO_MAX        = 0x8000
	_NGROUPS_MAX        = 0x10000
	_NSS_BUFLEN_GROUP   = 0x400
	_NSS_BUFLEN_PASSWD  = 0x400
	_OPEN_MAX           = 0x100
	_PTHREAD_KEYS_MAX   = 0x400
	_PTHREAD_STACK_MIN  = 0x4000
	_RE_DUP_MAX         = 0x7fff
	_RTSIG_MAX          = 0x20
	_SEM_VALUE_MAX      = 0x7fffffff
	_STREAM_MAX         = 0x10
	_SYMLOOP_MAX        = -0x1
	_TTY_NAME_MAX       = 0x20

	_UIO_MAXIOV = 0x400

	_INT_MAX = 0x7fffffff

	_POSIX_ADVISORY_INFO                = 0x31069
	_POSIX_ARG_MAX                      = 0x1000
	_POSIX_ASYNCHRONOUS_IO              = 0x31069
	_POSIX_BARRIERS                     = 0x31069
	_POSIX_CHILD_MAX                    = 0x19
	_POSIX_CLOCK_SELECTION              = 0x31069
	_POSIX_CPUTIME                      = 0x0
	_POSIX_FSYNC                        = 0x31069
	_POSIX_IPV6                         = 0x31069
	_POSIX_JOB_CONTROL                  = 0x1
	_POSIX_MAPPED_FILES                 = 0x31069
	_POSIX_MEMLOCK                      = 0x31069
	_POSIX_MEMLOCK_RANGE                = 0x31069
	_POSIX_MEMORY_PROTECTION            = 0x31069
	_POSIX_MESSAGE_PASSING              = 0x31069
	_POSIX_MONOTONIC_CLOCK              = 0x0
	_POSIX_PRIORITIZED_IO               = 0x31069
	_POSIX_PRIORITY_SCHEDULING          = 0x31069
	_POSIX_RAW_SOCKETS                  = 0x31069
	_POSIX_READER_WRITER_LOCKS          = 0x31069
	_POSIX_REALTIME_SIGNALS             = 0x31069
	_POSIX_REGEXP                       = 0x1
	_POSIX_SAVED_IDS                    = 0x1
	_POSIX_SEMAPHORES                   = 0x31069
	_POSIX_SHARED_MEMORY_OBJECTS        = 0x31069
	_POSIX_SHELL                        = 0x1
	_POSIX_SIGQUEUE_MAX                 = 0x20
	_POSIX_SPAWN                        = 0x31069
	_POSIX_SPIN_LOCKS                   = 0x31069
	_POSIX_SPORADIC_SERVER              = -0x1
	_POSIX_SYNCHRONIZED_IO              = 0x31069
	_POSIX_THREAD_ATTR_STACKADDR        = 0x31069
	_POSIX_THREAD_ATTR_STACKSIZE        = 0x31069
	_POSIX_THREAD_DESTRUCTOR_ITERATIONS = 0x4
	_POSIX_THREAD_PRIO_INHERIT          = 0x31069
	_POSIX_THREAD_PRIO_PROTECT          = 0x31069
	_POSIX_THREAD_PRIORITY_SCHEDULING   = 0x31069
	_POSIX_THREAD_PROCESS_SHARED        = 0x31069
	_POSIX_THREAD_SAFE_FUNCTIONS        = 0x31069
	_POSIX_THREAD_SPORADIC_SERVER       = -0x1
	_POSIX_THREADS                      = 0x31069
	_POSIX_TIMEOUTS                     = 0x31069
	_POSIX_TIMERS                       = 0x31069
	_POSIX_TRACE                        = -0x1
	_POSIX_TRACE_EVENT_FILTER           = -0x1
	_POSIX_TRACE_INHERIT                = -0x1
	_POSIX_TRACE_LOG                    = -0x1
	_POSIX_TYPED_MEMORY_OBJECTS         = -0x1
	_POSIX_VERSION                      = 0x31069

	_POSIX_V7_ILP32_OFF32  = 0x1
	_POSIX_V7_ILP32_OFFBIG = 0x1
	_POSIX_V7_LP64_OFF64   = -0x1
	_POSIX_V7_LPBIG_OFFBIG = -0x1

	_POSIX_V6_ILP32_OFF32  = 0x1
	_POSIX_V6_ILP32_OFFBIG = 0x1
	_POSIX_V6_LP64_OFF64   = -0x1
	_POSIX_V6_LPBIG_OFFBIG = -0x1

	_POSIX2_C_BIND    = 0x31069
	_POSIX2_C_DEV     = 0x31069
	_POSIX2_C_VERSION = 0x31069
	_POSIX2_CHAR_TERM = 0x31069
	_POSIX2_LOCALEDEF = 0x31069
	_POSIX2_SW_DEV    = 0x31069
	_POSIX2_VERSION   = 0x31069

	_XOPEN_ENH_I18N         = 0x1
	_XOPEN_REALTIME         = 0x1
	_XOPEN_REALTIME_THREADS = 0x1
	_XOPEN_SHM              = 0x1
	_XOPEN_UNIX             = 0x1
	_XOPEN_VERSION          = 0x2bc
	_XOPEN_XCU_VERSION      = 0x4
)
