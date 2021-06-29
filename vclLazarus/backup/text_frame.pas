unit text_frame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, StdCtrls;

type

  { TTextFrame }

  TTextFrame = class(TFrame)
    Memo: TMemo;
  private

  public

  end;

implementation

{$R *.lfm}

end.

