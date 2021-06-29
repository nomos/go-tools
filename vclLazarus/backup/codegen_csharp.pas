unit codegen_csharp;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, Buttons, StdCtrls;

type

  { TCodeGenCSharp }

  TCodeGenCSharp = class(TFrame)
    DistDirLabel: TEdit;
    ModelDirLabel: TEdit;
    GenerateButton: TButton;
    HelpButton: TSpeedButton;
    Label1: TLabel;
    Label2: TLabel;
    OpenDistDirButton: TSpeedButton;
    OpenModelDirButton: TSpeedButton;
    StaticText1: TStaticText;
  private

  public

  end;

var
  CodeGenCSharp: TCodeGenCSharp;

implementation

{$R *.lfm}

end.

