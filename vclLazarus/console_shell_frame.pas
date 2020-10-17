unit console_shell_frame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, ExtCtrls, StdCtrls, Buttons, ComCtrls;

type

  { TConsoleShell }

  TConsoleShell = class(TFrame)
    BottomPanel: TPanel;
    ShellSelect: TComboBox;
    Console: TMemo;
    Panel1: TPanel;
    SendButton: TButton;
    CmdEdit: TEdit;
    procedure CmdEditChange(Sender: TObject);
    procedure SendButtonClick(Sender: TObject);
  private

  public

  end;

implementation

{$R *.lfm}

{ TConsoleShell }

procedure TConsoleShell.CmdEditChange(Sender: TObject);
begin

end;

procedure TConsoleShell.SendButtonClick(Sender: TObject);
begin

end;

end.

