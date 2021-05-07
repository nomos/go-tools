unit model_gen;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, StdCtrls, Buttons;

type

  { TModelGenFrame }

  TModelGenFrame = class(TFrame)
    GenButton: TButton;
    Label14: TLabel;
    Label15: TLabel;
    Label16: TLabel;
    OpenCsFolderButton: TSpeedButton;
    OpenCsFolderButton1: TSpeedButton;
    OpenCsFolderLabel: TLabel;
    OpenCsFolderLabel1: TLabel;
    OpenModelFolderButton: TSpeedButton;
    OpenModelFolderLabel: TLabel;
    procedure Label16Click(Sender: TObject);
  private

  public

  end;

implementation

{$R *.lfm}

{ TModelGenFrame }

procedure TModelGenFrame.Label16Click(Sender: TObject);
begin

end;

end.

