unit imagecutter;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Buttons, StdCtrls;

type

  { TImageCutter }

  TImageCutter = class(TFrame)
    ImageHeight: TEdit;
    ImageWidth: TEdit;
    Label2: TLabel;
    Label3: TLabel;
    PngPath: TEdit;
    GenerateButton: TButton;
    Label1: TLabel;
    OpenPngButton: TSpeedButton;
  private

  public

  end;

implementation

{$R *.lfm}

end.

