from bokeh.plotting import figure, output_file, show

if __name__ == '__main__':
    output_file('simple_graphic.html')
    fig = figure()
    
    all_values = int(input('How to many values to graphic? '))
    
    x_values = list(range(all_values))
    y_values = []

    for i in x_values:
        value = int(input(f'How to valor for y {i}: '))
        y_values.append(value)

    fig.line(x_values, y_values)
    show(fig)
