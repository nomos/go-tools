unit chart_frame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, ExtCtrls;

type

  { TChartFrame }

  TChartFrame = class(TFrame)
    Main: TPaintBox;
    procedure MainClick(Sender: TObject);
  private

  public

  end;

implementation

{$R *.lfm}

{ TChartFrame }

procedure TChartFrame.MainClick(Sender: TObject);
begin

end;

end.

