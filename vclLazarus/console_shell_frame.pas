unit console_shell_frame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, ExtCtrls, StdCtrls, Buttons, ComCtrls;

type

  { TConsoleShell }

  TConsoleShell = class(TFrame)
    BottomPanel: TPanel;
    Memo1: TMemo;
    SendButton: TButton;
    CmdEdit: TEdit;
  private

  public

  end;

implementation

{$R *.lfm}

end.

