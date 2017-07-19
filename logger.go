package logger

import (
	"errors"
	"log"
	"os"
)

type Logger struct {
	Stage_error   *log.Logger
	Stage_info    *log.Logger
	Process_error *log.Logger
	Process_info  *log.Logger
	Log_dir       string
}

func (this *Logger) Init() error {
	if this.Log_dir == "" {
		err := errors.New("log dir is nil")
		return err
	}

	err := os.MkdirAll(this.Log_dir, os.ModeDir|0766)
	if err != nil {
		return err
	}

	stage_error_path := this.Log_dir + "/" + "StageError.log"
	stage_info_path := this.Log_dir + "/" + "StageInfo.log"
	process_error_path := this.Log_dir + "/" + "ProcessError.log"
	process_info_path := this.Log_dir + "/" + "ProcessInfo.log"

	f_stage_error, err := os.OpenFile(stage_error_path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	f_stage_info, err := os.OpenFile(stage_info_path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	f_process_error, err := os.OpenFile(process_error_path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	f_process_info, err := os.OpenFile(process_info_path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	this.Stage_error = log.New(f_stage_error, "StageError:", log.Llongfile|log.LstdFlags)
	this.Stage_info = log.New(f_stage_info, "StageInfo:", log.Llongfile|log.LstdFlags)
	this.Process_error = log.New(f_process_error, "ProcessError:", log.Llongfile|log.LstdFlags)
	this.Process_info = log.New(f_process_info, "ProcessInfo", log.Llongfile|log.LstdFlags)

	return nil
}

func (this *Logger) StageError(v interface{}) error {
	if this.Stage_error == nil {
		errors.New("Attribute Stage_error is nil")
	}
	this.Stage_error.Print(v)

	return nil
}

func (this *Logger) StageInfo(v interface{}) error {
	if this.Stage_info == nil {
		errors.New("Attribute Stage_info is nil")
	}
	this.Stage_info.Print(v)

	return nil
}

func (this *Logger) ProcessError(v interface{}) error {
	if this.Process_error == nil {
		errors.New("Attribute Process_error is nil")
	}
	this.Process_error.Print(v)

	return nil
}

func (this *Logger) ProcessInfo(v interface{}) error {
	if this.Process_info == nil {
		errors.New("Attribute Process_info is nil")
	}
	this.Process_info.Print(v)

	return nil
}
