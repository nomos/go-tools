unit icon_button;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Buttons;

type

  { TIconButton }

  TIconButton = class(TFrame)
    Button: TSpeedButton;
    procedure ButtonClick(Sender: TObject);
  private

  public

  end;

implementation

{$R *.lfm}

{ TIconButton }

procedure TIconButton.ButtonClick(Sender: TObject);
begin

end;

end.

