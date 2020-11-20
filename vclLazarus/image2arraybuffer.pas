unit image2arraybuffer;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, ExtCtrls, StdCtrls;

type

  { TImage2ArrayBuffer }

  TImage2ArrayBuffer = class(TFrame)
    Label1: TLabel;
    DropDownPanel: TPanel;
    procedure Label1Click(Sender: TObject);
  private

  public

  end;

implementation

{$R *.lfm}

{ TImage2ArrayBuffer }

procedure TImage2ArrayBuffer.Label1Click(Sender: TObject);
begin

end;

end.

