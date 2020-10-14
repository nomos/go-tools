unit console_shell_frame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, ExtCtrls, StdCtrls, Buttons, ComCtrls;

type

  { TConsoleShell }

  TConsoleShell = class(TFrame)
    BottomPanel: TPanel;
    Console: TMemo;
    Panel1: TPanel;
    SendButton: TButton;
    CmdEdit: TEdit;
  private

  public

  end;

implementation

{$R *.lfm}

end.

