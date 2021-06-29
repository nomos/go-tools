program vcltool;

{$mode objfpc}{$H+}

uses
  {$IFDEF UNIX}{$IFDEF UseCThreads}
  cthreads,
  {$ENDIF}{$ENDIF}
  Interfaces, // this includes the LCL widgetset
  Forms, auto_deploy_frame, form_list_frame, deploy_frame, excel2jsonminigame,
  image2arraybuffer, imagecutter, empty, model_gen, codegen_csharp;

{$R *.res}

begin
  RequireDerivedFormResource:=True;
  Application.Scaled:=True;
  Application.Initialize;
  Application.CreateForm(TEmptyForm, EmptyForm);
  Application.CreateForm(TCodeGenCSharp, CodeGenCSharp);
  Application.Run;
end.

