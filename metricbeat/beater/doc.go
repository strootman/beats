/*
Package beater provides the implementation of the libbeat Beater interface for
Metricbeat and functions for running Metricbeat Modules on their own.

Metricbeat collects metrics from operating systems and services. The code for
gathering metrics from a particular service is organized into a logical grouping
called a Module. Each Module has one or more MetricSet implementations which
do the work of collecting a particular set of metrics from the service.

The public interfaces used in implementing Modules and MetricSets are defined in
the github.com/elastic/beats/metricbeat/mb package.

Event Format

Each event generated by Metricbeat has the same general structure. The example
event below was generated by a MetricSet named "cpu" in the "system" Module.

	{
	  "@timestamp": "2016-05-11T03:36:51.518Z",
	  "beat": {
		"hostname": "host.example.com",
		"name": "host.example.com"
	  },
	  "metricset": "cpu",
	  "module": "system",
	  "rtt": 1783,
	  "system-cpu": {
		"idle": 211609484,
		"iowait": 8244,
		"irq": 424,
		"nice": 0,
		"softirq": 26458,
		"steal": 0,
		"system": 792081,
		"system_p": 0,
		"user": 1677782,
		"user_p": 0
	  },
	  "type": "metricsets"
	}

All events are stored in one index called metricbeat by default. Each
MetricSet's data format is potentially unique so the MetricSet data is added to
event as a dictionary under a key that is unique to the MetricSet. The key
is constructed from the Module name and MetricSet name to ensure uniqueness.
All documents are stored under the same type called "metricsets".
*/
package beater